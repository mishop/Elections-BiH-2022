from benfordslaw import benfordslaw

# Initialize
bl = benfordslaw(pos=-1)

# Load elections example
df = bl.import_example(data='RS')

# Extract election information.
X = df['totalVotes'].loc[df['name']=='Kandidat'].values

# Make fit
results = bl.fit(X)

# Plot
bl.plot(title='Test prevare zadnje cifre RS - Kandidat by Misho', barcolor=[0.5, 0.5, 0.5], fontsize=12, barwidth=0.4)
